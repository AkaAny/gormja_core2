import {dbField, DBType} from "../sdk/db_field";
import {SourceEntityTrait} from "../sdk/db";

export class JLFutureTargetItem implements SourceEntityTrait{
    @dbField("id",DBType.int64)
    ID:number;
    @dbField("school_code",DBType.string)
    SchoolCode:string;
    @dbField("staff_id",DBType.string)
    StaffID:string;
    @dbField("future_target",DBType.string)
    FutureTarget:string;

    constructor(props:any) {
        this.ID=props.ID;
        this.SchoolCode=props.SchoolCode;
        this.StaffID=props.StaffID;
        this.FutureTarget=props.FutureTarget;
    }

    static newModel():JLFutureTargetItem{
        return new JLFutureTargetItem({});
    }

    static tableName():string{
        return "future_target_items"
    }
}