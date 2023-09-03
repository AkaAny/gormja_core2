import {dbField, DBType} from "../sdk/db_field";

export interface GradeOverviewTrait{
    SchoolCode:string;
    Grade:string;
    GraduateStudentCount:number;
}

export class GradeOverview implements GradeOverviewTrait{

    @dbField("school_code",DBType.string)
    SchoolCode:string;

    @dbField("grade",DBType.string)
    Grade:string;

    @dbField("graduate_student_count",DBType.int64)
    GraduateStudentCount:number;

    constructor(props:GradeOverviewTrait) {
        this.SchoolCode=props.SchoolCode;
        this.Grade=props.Grade;
        this.GraduateStudentCount=props.GraduateStudentCount;
    }

}