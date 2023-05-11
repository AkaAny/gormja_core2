import {dbField, DBType} from "./sdk/db_field";
import {SourceEntityTrait} from "./sdk/db";

export class HDUStudentDetail implements SourceEntityTrait{
    @dbField("STAFF_ID",DBType.string)
    StaffID:string;
    @dbField("STAFF_NAME",DBType.string)
    StaffName:string;
    @dbField("GENDER",DBType.string)
    Gender:string;
    @dbField("CLASS_ID",DBType.string)
    ClassID:string;
    @dbField("GRADE",DBType.string)
    Grade:string;
    @dbField("UNIT_ID",DBType.string)
    UnitCode:string;
    @dbField("MAJOR_CODE",DBType.string)
    MajorCode:string;

    constructor(props:any) {
        this.StaffID=props.StaffID;
        this.StaffName=props.StaffName;
        this.Gender=props.Gender;
        this.Grade=props.Grade;
        this.ClassID=props.ClassID;
        this.UnitCode=props.UnitCode;
        this.MajorCode=props.MajorCode;
    }

    static newModel():HDUStudentDetail{
        return new HDUStudentDetail({});
    }

    static tableName():string{
        return "HDUHELP_VIEW_STUDENT_DETAIL"
    }
}