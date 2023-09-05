import {dbField, DBType} from "../../sdk/db_field";
import {SourceEntityTrait} from "../../sdk/db";

export class HDUClassInfo implements SourceEntityTrait{

    @dbField("GRADE",DBType.string)
    Grade:string;

    @dbField("CLASSID",DBType.string)
    ClassID:string;

    @dbField("FDY_STAFFID",DBType.string)
    CounselorStaffID:string;

    constructor(props:any) {
        this.Grade=props.Grade;
        this.ClassID=props.ClassID;
        this.CounselorStaffID=props.CounselorStaffID;
    }

    static tableName(): string {
        return "JL_VIEW_CLASS_INFO";
    }

}