import {dbField, DBType} from "./sdk/db_field";
import {SourceEntityTrait} from "./sdk/db";

export class HDUStudentDetail implements SourceEntityTrait{
    @dbField("STAFFID",DBType.string)
    StaffID:string;
    @dbField("STAFFNAME",DBType.string)
    StaffName:string;
    @dbField("GENDER",DBType.string)
    Gender:string;
    @dbField("CLASSID",DBType.string)
    ClassID:string;
    @dbField("GRADE",DBType.string)
    Grade:string;
    @dbField("UNITCODE",DBType.string)
    UnitCode:string;
    @dbField("UNITNAME",DBType.string)
    UnitName:string;
    @dbField("MAJORCODE",DBType.string)
    MajorCode:string;
    @dbField("MAJORNAME",DBType.string)
    MajorName:string;
    @dbField("TEACHERID",DBType.string)
    CounselorStaffID:string;
    @dbField("TEACHERNAME",DBType.string)
    CounselorStaffName:string;

    constructor(props:any) {
        this.StaffID=props.StaffID;
        this.StaffName=props.StaffName;
        this.Gender=props.Gender;
        this.Grade=props.Grade;
        this.ClassID=props.ClassID;
        this.UnitCode=props.UnitCode;
        this.UnitName=props.UnitName;
        this.MajorCode=props.MajorCode;
        this.MajorName=props.MajorName;
        this.CounselorStaffID=props.CounselorStaffID;
        this.CounselorStaffName=props.CounselorStaffName;
    }

    static newModel():HDUStudentDetail{
        return new HDUStudentDetail({});
    }

    static tableName():string{
        return "HDUHELP_VIEW_STUDENT_INFO"
    }
}