import {DBService, ServiceTrait} from "./sdk/base_service";
import {Score, ScoreTrait} from "./unify/score";
import {HDUScore, HDUScoreTypeNeedToSkip} from "./hdu_score";
import {log} from "./sdk/logger";
import {Day} from "./sdk/duration";

interface ScoreLookupTrait{
    StaffID:string;
    SchoolYear?:string;
    Semester?:string;
}

export class ScoreService extends DBService implements ServiceTrait<Score,ScoreLookupTrait>{
    constructor() {
        super("score",{
            dataSourceID:"hdu-oracle",
        },{
            ttl: 1*Day,
        });
    }

    lookup(conds: ScoreLookupTrait): Score[] {
        let tx=this.getDB().startSession(HDUScore)
            .where("STAFFID=?",conds.StaffID);
        if (conds.SchoolYear){
            tx=tx.where("SCHOOLYEAR=?",conds.SchoolYear);
        }
        if(conds.Semester){
            tx=tx.where("SEMESTER=?",conds.Semester);
        }
        const scoreItems= tx.find() as HDUScore[];
        //getRuntime().debugBreakpoint(this,scoreItems);
        const unifyScoreItems= scoreItems.filter((item)=>{
            return item.Score !== HDUScoreTypeNeedToSkip;
        }).map((item)=>{
            const creditNumber=Number(item.Credit);
            let scoreNumber=Number(item.Score);
            if(isNaN(scoreNumber)){
                switch (item.Score) {
                    case "优秀":
                        scoreNumber=100;
                        break;
                    case "良好":
                        scoreNumber=80;
                        break;
                    case "中等":
                        scoreNumber=60;
                        break;
                    case  "合格":
                        scoreNumber=40;
                        break;
                    case "及格":
                        scoreNumber=40;
                        break;
                    case "不及格":
                        scoreNumber=20;
                        break;
                    default:
                        throw `no role to handle score item:${item}`;
                }
            }
            if(scoreNumber<60 && item.ScoreBK){
                scoreNumber=item.ScoreBK;
            }
            return new Score({
                SchoolCode: 'hdu',
                StaffID: item.StaffID,
                CourseCode: item.CourseCode,

                CourseName: item.CourseName,
                SchoolYear: item.SchoolYear,
                Semester: item.Semester,

                Credit:creditNumber,
                Score:scoreNumber,
            })
        });
        log(unifyScoreItems);
        return unifyScoreItems;
    }

    newUnifyModel():Score{
        return new Score({} as ScoreTrait);
    }
}