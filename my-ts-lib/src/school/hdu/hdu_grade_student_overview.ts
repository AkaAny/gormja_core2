import {dbField, DBType} from "../../sdk/db_field";
import {SourceEntityTrait} from "../../sdk/db";

export class HDUGradeOverview implements SourceEntityTrait{

    @dbField("GRADUATE_STUDENT_COUNT",DBType.int64)
    GraduateStudentCount:number;

    @dbField("GRADE",DBType.string)
    Grade:string;

    constructor(props:any) {
        this.GraduateStudentCount=props.GraduateStudentCount;
        this.Grade=props.Grade;
    }

    static tableName(): string {
        return "JL_VIEW_GRADUATE";
    }

}