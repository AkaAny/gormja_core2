import {dbField, DBType} from "../sdk/db_field";

export interface JLUserRegisterOverviewTrait{
    SchoolCode:string;
    UserType:string;
    Grade?:string;
    UserRegisterCount:number;
}

export class JLUserRegisterOverview implements JLUserRegisterOverviewTrait{

    @dbField("school_code",DBType.string)
    SchoolCode:string;

    @dbField("user_type",DBType.string)
    UserType:string;

    @dbField("grade",DBType.string)
    Grade?:string;

    @dbField("user_register_count",DBType.int64)
    UserRegisterCount:number;

    constructor(props:JLUserRegisterOverviewTrait) {
        this.SchoolCode=props.SchoolCode;
        this.UserType=props.UserType;
        this.Grade=props.Grade;
        this.UserRegisterCount=props.UserRegisterCount;
    }

}