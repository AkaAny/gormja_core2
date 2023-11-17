import {DBService, ServiceTrait} from "./sdk/base_service";
import {Rank, RankTrait} from "./unify/rank";
import {log} from "./sdk/logger";
import {getRuntime} from "./sdk/runtime";
import {ProfileService} from "./profile_service";
import {ScoreService} from "./score_service";

interface StaffIDGPATrait{
    staffID:string;
    gpa:number;
}

interface RankLookupTrait{
    StaffID:string;
}

export class RankService extends DBService implements ServiceTrait<Rank,RankLookupTrait>{
    constructor() {
        super("rank",{
            dataSourceID:'hdu-oracle',
        },false);
    }

    lookup(conds:RankLookupTrait):Rank[]{
        log(conds);
        const profileService=getRuntime().getService(ProfileService);
        const staffIDMatched= profileService.lookup({
            StaffID:conds.StaffID,
        });
        log("matched staff ids:",staffIDMatched);
        if(staffIDMatched.length===0){
            throw `staff id:${conds.StaffID} does not exist`;
        }
        const selfProfile=staffIDMatched[0];
        const studentDetails=profileService.lookupByGradeAndMajor({
            Grade:selfProfile.Grade,
            UnitCode: selfProfile.UnitCode,
            MajorCode: selfProfile.MajorCode,
        })
        // const studentDetails=this.db.startSession(HDUStudentDetail)
        //     .where("GRADE = ?",`${conds.Grade}`)
        //     .where("UNIT_ID = ?",conds.UnitCode)
        //     .where("MAJOR_CODE=?",conds.MajorCode)
        //     .find() as HDUStudentDetail[];
        log(studentDetails);
        //lookup score, actually another service should be wrapped by runtime to add cache
        const scoreService=getRuntime().getService(ScoreService);
        const staffIDGPAs= studentDetails.map((studentItem)=>{
            const scoreItems= scoreService.lookup({
                StaffID: studentItem.StaffID,
            });

            getRuntime().debugBreakpoint("after score service lookup",this,scoreItems);
            let sumScoreWithCredit=0;
            let sumCredit=0;
            scoreItems.forEach((item)=>{
                sumCredit+=item.Credit;
                sumScoreWithCredit+=(item.Score*item.Credit);
                return;
            })
            let gpa=sumScoreWithCredit/sumCredit;
            const staffIDGPAPair={
                staffID:studentItem.StaffID,
                gpa:gpa,
            } as StaffIDGPATrait;
            log("pair:",staffIDGPAPair)
            return staffIDGPAPair;
        });
        let gpaItem:StaffIDGPATrait= {
            staffID:conds.StaffID,
            gpa:0.0,
        };
        let rank= 0;
        staffIDGPAs.sort((a:StaffIDGPATrait,b:StaffIDGPATrait)=>{
            return b.gpa-a.gpa;
        }).forEach((item,index)=>{
            if(item.staffID!==conds.StaffID){
                return;
            }
            gpaItem=item;
            rank=index+1;
        })

        const result=new Rank({
            SchoolCode:'hdu',
            StaffID:conds.StaffID,
            GPA:gpaItem.gpa,
            Rank:rank,
            RankPercent:rank/studentDetails.length,
        });
        log(result);
        return [result];
    }

    newUnifyModel():Rank{
        return new Rank({} as RankTrait);
    }
}