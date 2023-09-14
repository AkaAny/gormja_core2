import {dbField, DBType} from "../sdk/db_field";

export interface JLInternshipExperienceTrait{
    ID:number;
    SchoolCode:string;
    StaffID:string;
    CompanyName:string;
    JobName:string;
    StartAt:Date;
    EndAt:Date;
    Description:string;
}

export class JLInternshipExperience implements JLInternshipExperienceTrait{
    @dbField("id",DBType.int64)
    ID:number;
    @dbField("school_code",DBType.string)
    SchoolCode:string;
    @dbField("staff_id",DBType.string)
    StaffID:string;
    @dbField("company_name",DBType.string)
    CompanyName:string;
    @dbField("job_name",DBType.string)
    JobName:string;
    @dbField("start_at",DBType.dateTime)
    StartAt:Date;
    @dbField("end_at",DBType.dateTime)
    EndAt:Date;
    @dbField("description",DBType.string)
    Description:string;

    constructor(props:JLInternshipExperienceTrait) {
        this.ID=props.ID;
        this.SchoolCode=props.SchoolCode;
        this.StaffID=props.StaffID;
        this.CompanyName=props.CompanyName;
        this.JobName=props.JobName;
        this.StartAt=props.StartAt;
        this.EndAt=props.EndAt;
        this.Description=props.Description;
    }
}