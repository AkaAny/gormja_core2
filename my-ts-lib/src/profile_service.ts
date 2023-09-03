import {DBService, ServiceTrait} from "./sdk/base_service";
import {Day} from "./sdk/duration";
import {Profile, ProfileTrait} from "./unify/profile";
import {log} from "./sdk/logger";
import {HDUStudentDetail} from "./school/hdu/hdu_student_detail";

interface PersonLookupTrait{
    //SchoolCode:string;
    StaffID?:string;
    ClassIDs?:string[];
}

export class ProfileService extends DBService implements ServiceTrait<Profile, PersonLookupTrait>{
    //ttl with manual refresh api
    constructor() {
        super("profile",{
            dataSourceID:'hdu-oracle',
        },{
            ttl:1*Day,
        });
    }

    lookup(params: PersonLookupTrait): Profile[] {
        let tx=this.getDB().startSession(HDUStudentDetail);
        log(params);
        if(params.StaffID){
            tx=tx.where("STAFFID=?",params.StaffID);
        }else{
            if(params.ClassIDs && Array.isArray(params.ClassIDs) && params.ClassIDs.length>0){
                tx=tx.where("CLASSID IN ?",params.ClassIDs);
            }else{
                return [];
            }
        }
        const studentDetails= tx.find() as HDUStudentDetail[];
        const items=studentDetails.map(this.toUnify);
        log(items);
        return items;
    }

    toUnify(item:HDUStudentDetail):Profile{
        return new Profile({
            Grade: item.Grade,
            MajorCode: item.MajorCode,
            MajorName:item.MajorName,
            SchoolCode: "hdu",
            StaffID: item.StaffID,
            StaffName:item.StaffName,
            Gender:item.Gender=="1"?"男":"女",
            ClassID:item.ClassID,
            UnitCode: item.UnitCode,
            UnitName:item.UnitName,
            CounselorStaffID: item.CounselorStaffID,
            CounselorStaffName: item.CounselorStaffName,
        });
    }

    lookupByGradeAndMajor(conds:{
        Grade:string,
        UnitCode:string,
        MajorCode:string,
    }):Profile[]{
        let studentDetails=this.getDB().startSession(HDUStudentDetail)
            .where("GRADE = ?",`${conds.Grade}`)
            .where("UNIT_ID = ?",conds.UnitCode)
            .where("MAJOR_CODE=?",conds.MajorCode)
            .find() as HDUStudentDetail[];
        const items=studentDetails.map(this.toUnify);
        log(items);
        return items;
    }

    newUnifyModel():Profile{
        return new Profile({} as ProfileTrait);
    }


}