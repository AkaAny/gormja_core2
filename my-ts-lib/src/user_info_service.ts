import {DBService, ServiceTrait} from "./sdk/base_service";
import {Day} from "./sdk/duration";

import {log} from "./sdk/logger";
import {HDUPersonInfo} from "./school/hdu/hdu_person_info";
import {stringifyUserType, UserInfo, UserInfoTrait} from "./unify/user_info";

interface UserInfoLookupTrait{
    //SchoolCode:string;
    StaffID:string;
    //ClassIDs?:string[];
}

export class UserInfoService extends DBService implements ServiceTrait<UserInfo, UserInfoLookupTrait>{
    //ttl with manual refresh api
    constructor() {
        super("user-info",{
            dataSourceID:'hdu-oracle',
        },{
            ttl:1*Day,
        });
    }

    lookup(params: UserInfoLookupTrait): UserInfo[] {
        let tx=this.getDB().startSession(HDUPersonInfo);
        log(params);
        tx=tx.where("STAFFID=?",params.StaffID);
        const personInfos= tx.find() as HDUPersonInfo[];
        const items=personInfos.map(this.toUnify);
        log(items);
        return items;
    }

    toUnify(item:HDUPersonInfo):UserInfo{
        return new UserInfo({
            DeployID: "hdu",
            UserID: item.StaffID,
            UserName:item.StaffName,
            UserType:stringifyUserType(item.StaffType),
        });
    }

    newUnifyModel():UserInfo{
        return new UserInfo({} as UserInfoTrait);
    }

}