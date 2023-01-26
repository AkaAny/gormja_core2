import {DBService, ServiceTrait} from "./sdk/base_service";
import {Day} from "./sdk/duration";
import {Profile, ProfileTrait} from "./unify/profile";
import {log} from "./sdk/logger";
import {HDUStudentDetail} from "./hdu_student_detail";

interface PersonLookupTrait{
    //SchoolCode:string;
    StaffID:string,
}

export class ProfileService extends DBService implements ServiceTrait<Profile, PersonLookupTrait>{
    //ttl with manual refresh api
    constructor() {
        super({
            dataSourceID:'hdu-oracle',
        },{
            ttl:1*Day,
        });
    }

    lookup(params: PersonLookupTrait): Profile[] {
        let tx=this.getDB().startSession(HDUStudentDetail)
            .where("STAFF_ID=?",params.StaffID);
        const studentDetails= tx.find() as HDUStudentDetail[];
        const items=studentDetails.map(this.toUnify);
        log(items);
        return items;
    }

    toUnify(item:HDUStudentDetail):Profile{
        return new Profile({
            Grade: item.Grade,
            MajorCode: item.MajorCode,
            SchoolCode: "hdu",
            StaffID: item.StaffID,
            UnitCode: item.UnitCode,
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