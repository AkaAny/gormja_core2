import {SourceEntityTrait} from "../sdk/db";
import {dbField, DBType} from "../sdk/db_field";

export class JLRecommendForStudentItem implements SourceEntityTrait{
    @dbField("student_id",DBType.string)
    StaffID:string;
    @dbField("job_id",DBType.int64)
    JobID:number;
    @dbField("score",DBType.float64)
    Score:number;

    constructor(props:any) {
        this.StaffID=props.StaffID;
        this.JobID=props.JobID;
        this.Score=props.Score;
    }

    static newModel():JLRecommendForStudentItem{
        return new JLRecommendForStudentItem({});
    }

    static tableName():string{
        return "recommend_for_students"
    }
}