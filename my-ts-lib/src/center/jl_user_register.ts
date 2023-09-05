import {dbField, DBType} from "../sdk/db_field";
import {SourceEntityTrait} from "../sdk/db";

export class JLUserRegisterItem implements SourceEntityTrait{
    @dbField("id",DBType.int64)
    ID:number;
    @dbField("school_code",DBType.string)
    SchoolCode:string;
    @dbField("staff_id",DBType.string)
    StaffID:string;
    @dbField("user_type",DBType.string)
    UserType:string;
    @dbField("grade",DBType.string)
    Grade:string;

    constructor(props:any) {
        this.ID=props.ID;
        this.SchoolCode=props.SchoolCode;
        this.StaffID=props.StaffID;
        this.UserType=props.UserType;
        this.Grade=props.Grade;
    }

    static newModel():JLUserRegisterItem{
        return new JLUserRegisterItem({});
    }

    static tableName():string{
        return "user_register_items"
    }
}