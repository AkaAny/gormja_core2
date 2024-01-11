import {dbField, DBType} from "../../sdk/db_field";
import {SourceEntityTrait} from "../../sdk/db";

export class HDUPersonInfo implements SourceEntityTrait{
    @dbField("STAFFID",DBType.string)
    StaffID:string;
    @dbField("STAFFNAME",DBType.string)
    StaffName:string;
    @dbField("STAFFTYPE",DBType.string)
    StaffType:string;


    constructor(props:any) {
        this.StaffID=props.StaffID;
        this.StaffName=props.StaffName;
        this.StaffType=props.StaffType;
    }

    static newModel():HDUPersonInfo{
        return new HDUPersonInfo({});
    }

    static tableName():string{
        return "HDUHELP_VIEW_PERSON_INFO"
    }
}