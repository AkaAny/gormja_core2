import {dbField, DBType} from "../sdk/db_field";

export interface ClassInfoTrait{
    SchoolCode:string;
    Grade:string;
    ClassID:string;
    CounselorStaffID:string;
}

export class ClassInfo implements ClassInfoTrait{
    @dbField("school_code",DBType.string)
    SchoolCode:string;

    @dbField("grade",DBType.string)
    Grade:string;

    @dbField("class_id",DBType.string)
    ClassID:string;

    @dbField("counselor_staff_id",DBType.string)
    CounselorStaffID:string;

    constructor(props:ClassInfoTrait) {
        this.SchoolCode=props.SchoolCode;
        this.Grade=props.Grade;
        this.ClassID=props.ClassID;
        this.CounselorStaffID=props.CounselorStaffID;
    }

}