import {DBService, ServiceTrait} from "./sdk/base_service";
import {Hour} from "./sdk/duration";
import {log} from "./sdk/logger";
import {JLUserRegisterItem} from "./center/jl_user_register";
import {JLUserRegisterOverview, JLUserRegisterOverviewTrait} from "./unify/jl_user_register_overview";

interface JLUserRegisterOverviewLookupTrait{
    SchoolCode:string;
    UserType:string;
    Grade?:string;
}

export class JLUserRegisterOverviewService extends DBService implements ServiceTrait<JLUserRegisterOverview, JLUserRegisterOverviewLookupTrait>{
    //ttl with manual refresh api
    constructor() {
        super("jl-user-register-overview",{
            dataSourceID:'hdu-jl-student-attribute',
        },{
            ttl:1*Hour,
        });
    }

    lookup(params: JLUserRegisterOverviewLookupTrait): JLUserRegisterOverview[] {
        let tx=this.getDB().startSession(JLUserRegisterItem)
            .where("school_code=?",params.SchoolCode);
        if(params.UserType){
            tx=tx.where("user_type=?",params.UserType);
        }
        if(params.Grade){
            tx=tx.where("grade=?",params.Grade);
        }
        log(params);
        const registerCount= tx.count();
        const result=new JLUserRegisterOverview({
            SchoolCode: params.SchoolCode,
            UserType:params.UserType,
            Grade: params.Grade,
            UserRegisterCount: registerCount,
        })
        log(result);
        return [result];
    }

    newUnifyModel():JLUserRegisterOverview{
        return new JLUserRegisterOverview({} as JLUserRegisterOverviewTrait);
    }


}