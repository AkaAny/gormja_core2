import {dbField, DBType} from "../sdk/db_field";
import {log} from "../sdk/logger";

export interface ScoreTrait{
    SchoolCode:string;
    StaffID:string;
    CourseCode:string;
    SchoolYear:string;
    Semester:string;
    CourseName:string;
    Credit:number;
    Score:number;
}

export class Score implements ScoreTrait{
    @dbField("school_code",DBType.string,{
        isPrimaryKey:true,
    })
    SchoolCode:string;

    @dbField("staff_id",DBType.string,{
        isPrimaryKey:true,
    })
    StaffID:string;

    @dbField("course_code",DBType.string,{
        isPrimaryKey:true,
    })
    CourseCode:string;

    @dbField("school_year",DBType.string)
    SchoolYear:string;

    @dbField("semester",DBType.string)
    Semester:string;

    @dbField("course_name",DBType.string)
    CourseName:string;

    @dbField("credit",DBType.float64)
    Credit:number;

    @dbField("score",DBType.float64)
    Score:number;

    constructor(props:ScoreTrait) {
        this.SchoolCode=props.SchoolCode;
        this.StaffID=props.StaffID;
        this.CourseCode=props.CourseCode;
        this.SchoolYear=props.SchoolYear;
        this.Semester=props.Semester;
        this.CourseName=props.CourseName;
        this.Credit=props.Credit;
        this.Score=props.Score;
    }

    static tableName():string{
        return "scores";
    }
}

log(Reflect.ownKeys(Score.prototype));