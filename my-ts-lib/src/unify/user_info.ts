import {dbField, DBType} from "../sdk/db_field";

export interface UserInfoTrait{
    DeployID:string;
    UserID:string;
    UserName:string;
    UserType:string;
}

export class UserInfo implements UserInfoTrait{
    @dbField("deploy_id",DBType.string,{
        isPrimaryKey:true,
    })
    DeployID:string;
    @dbField("user_id",DBType.string,{
        isPrimaryKey:true,
    })
    UserID:string;
    @dbField("user_name",DBType.string)
    UserName:string;
    @dbField("user_type",DBType.string)
    UserType:string;

    
    constructor(props:UserInfoTrait) {
        this.DeployID=props.DeployID;
        this.UserID=props.UserID;
        this.UserName=props.UserName;
        this.UserType=props.UserType;
    }
}

export function stringifyUserType(userTypeCode:string):string{
    switch (userTypeCode){
        case "1":
            return "undergraduate";
        case "2":
            return "teacher";
        case "3":
            return "postgraduate";
        default:
            return "unknown";
    }
}