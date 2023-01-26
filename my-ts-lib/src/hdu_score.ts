import {dbField, DBType} from "./sdk/db_field";
import {SourceEntityTrait} from "./sdk/db";

export const HDUScoreTypeNeedToSkip ='缓考'
export type HDUScoreStrType='优秀'| '良好' | '中等' | '合格' | '及格' | '不及格' | '缓考';
export type HDUScoreType=number | HDUScoreStrType;

export class HDUScore implements SourceEntityTrait{
    @dbField("STAFFID",DBType.string)
    StaffID:string;
    @dbField("SCHOOLYEAR",DBType.string)
    SchoolYear:string;
    @dbField("SEMESTER",DBType.string)
    Semester:string;
    @dbField("COURSECODE",DBType.string)
    CourseCode:string;
    @dbField("COURSE",DBType.string)
    CourseName:string;
    @dbField("CREDIT",DBType.string)
    Credit:string;
    @dbField("SCORE",DBType.string)
    Score: HDUScoreType;
    @dbField("SCORE_BK",DBType.float64)
    ScoreBK:number;

    constructor(props:any) {
        this.StaffID=props.StaffID;
        this.SchoolYear=props.SchoolYear;
        this.Semester=props.Semester;
        this.CourseCode=props.CourseCode;
        this.CourseName=props.CourseName;
        this.Credit=props.Credit;
        this.Score=props.Score;
        this.ScoreBK=props.ScoreBK;
    }

    static tableName(): string {
        return "HDUHELP_VIEW_STUDENT_GRADE";
    }

}