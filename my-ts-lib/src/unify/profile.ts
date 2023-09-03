import {dbField, DBType} from "../sdk/db_field";

export interface ProfileTrait{
    SchoolCode:string;
    StaffID:string;
    StaffName:string;
    Gender:string;
    ClassID:string;
    Grade:string;
    UnitCode:string;
    UnitName:string;
    MajorCode:string;
    MajorName:string;
    CounselorStaffID:string;
    CounselorStaffName:string;
}

export class Profile implements ProfileTrait{
    @dbField("school_code",DBType.string,{
        isPrimaryKey:true,
    })
    SchoolCode:string;
    @dbField("staff_id",DBType.string,{
        isPrimaryKey:true,
    })
    StaffID:string;
    @dbField("staff_name",DBType.string)
    StaffName:string;
    @dbField("gender",DBType.string)
    Gender:string;
    @dbField("class_id",DBType.string)
    ClassID:string;

    @dbField("grade",DBType.string)
    Grade:string;
    @dbField("unit_code",DBType.string)
    UnitCode:string;
    @dbField("unit_name",DBType.string)
    UnitName:string;
    @dbField("major_code",DBType.string)
    MajorCode:string;
    @dbField("major_name",DBType.string)
    MajorName:string;

    @dbField("counselor_staff_id",DBType.string)
    CounselorStaffID:string;
    @dbField("counselor_staff_name",DBType.string)
    CounselorStaffName:string;


    constructor(props:ProfileTrait) {
        this.SchoolCode=props.SchoolCode;
        this.StaffID=props.StaffID;
        this.StaffName=props.StaffName;
        this.Gender=props.Gender;
        this.ClassID=props.ClassID;
        this.Grade=props.Grade;
        this.UnitCode=props.UnitCode;
        this.UnitName=props.UnitName;
        this.MajorCode=props.MajorCode;
        this.MajorName=props.MajorName;
        this.CounselorStaffID=props.CounselorStaffID;
        this.CounselorStaffName=props.CounselorStaffName;
    }

}