import {dbField, DBType} from "../sdk/db_field";
import {SourceEntityTrait} from "../sdk/db";

export class JLInternshipExperienceItem implements SourceEntityTrait{
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

    constructor(props:any) {
        this.ID=props.ID;
        this.SchoolCode=props.SchoolCode;
        this.StaffID=props.StaffID;
        this.CompanyName=props.CompanyName;
        this.JobName=props.JobName;
        this.StartAt=props.StartAt;
        this.EndAt=props.EndAt;
        this.Description=props.Description;
    }

    static newModel():JLInternshipExperienceItem{
        return new JLInternshipExperienceItem({});
    }

    static tableName():string{
        return "internship_experience_items"
    }
}