import {dbField, DBType} from "../sdk/db_field";

export interface JLRecommendForStudentTrait{
    SchoolCode:string;
    StaffID:string;
    JobID:number;
    Score:number;
}

export class JLRecommendForStudent implements JLRecommendForStudentTrait{

    @dbField("school_code",DBType.string)
    SchoolCode:string;

    @dbField("staff_id",DBType.string)
    StaffID:string;

    @dbField("job_id",DBType.int64)
    JobID:number;

    @dbField("score",DBType.float64)
    Score:number;

    constructor(props:JLRecommendForStudentTrait) {
        this.SchoolCode=props.SchoolCode;
        this.StaffID=props.StaffID;
        this.JobID=props.JobID;
        this.Score=props.Score;
    }

}