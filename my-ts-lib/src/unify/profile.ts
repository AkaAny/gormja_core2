import {dbField, DBType} from "../sdk/db_field";

export interface ProfileTrait{
    SchoolCode:string;
    StaffID:string;
    StaffName:string;
    Gender:string;
    ClassID:string;
    Grade:string;
    UnitCode:string;
    MajorCode:string;
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
    @dbField("major_code",DBType.string)
    MajorCode:string;

    constructor(props:ProfileTrait) {
        this.SchoolCode=props.SchoolCode;
        this.StaffID=props.StaffID;
        this.StaffName=props.StaffName;
        this.Gender=props.Gender;
        this.ClassID=props.ClassID;
        this.Grade=props.Grade;
        this.UnitCode=props.UnitCode;
        this.MajorCode=props.MajorCode;
    }

}