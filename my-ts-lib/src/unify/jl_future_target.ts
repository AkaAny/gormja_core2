import {dbField, DBType} from "../sdk/db_field";

export interface JLFutureTargetTrait{
    ID:number;
    SchoolCode:string;
    StaffID:string;
    FutureTarget:string;
}

export class JLFutureTarget implements JLFutureTargetTrait{
    @dbField("id",DBType.int64)
    ID:number;

    @dbField("school_code",DBType.string)
    SchoolCode:string;

    @dbField("staff_id",DBType.string)
    StaffID:string;

    @dbField("future_target",DBType.string)
    FutureTarget:string;

    constructor(props:JLFutureTargetTrait) {
        this.ID=props.ID;
        this.SchoolCode=props.SchoolCode;
        this.StaffID=props.StaffID;
        this.FutureTarget=props.FutureTarget;
    }

}