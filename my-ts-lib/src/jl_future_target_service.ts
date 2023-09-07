import {DBService, ServiceTrait} from "./sdk/base_service";
import {Hour} from "./sdk/duration";
import {log} from "./sdk/logger";
import {JLFutureTarget, JLFutureTargetTrait} from "./unify/jl_future_target";
import {JLFutureTargetItem} from "./center/jl_future_target_item";

interface JLFutureTargetLookupTrait{
    SchoolCode:string;
    StaffIDs:string[];
}

export class JLFutureTargetService extends DBService implements ServiceTrait<JLFutureTarget, JLFutureTargetLookupTrait>{
    //ttl with manual refresh api
    constructor() {
        super("jl-future-target",{
            dataSourceID:'hdu-jl-student-attribute',
        },{
            ttl:1*Hour,
        });
    }

    lookup(params: JLFutureTargetLookupTrait): JLFutureTarget[] {
        let tx=this.getDB().startSession(JLFutureTargetItem)
            .where("school_code=?",params.SchoolCode)
            .where("staff_id IN ?",params.StaffIDs);
        log(params);
        const items= tx.find() as JLFutureTarget[];
        const results=items.map((item)=>{
            return new JLFutureTarget({
                ID: item.ID,
                SchoolCode: item.SchoolCode,
                StaffID: item.StaffID,
                FutureTarget: item.FutureTarget,
            })
        })
        log(results);
        return results;
    }

    newUnifyModel():JLFutureTarget{
        return new JLFutureTarget({} as JLFutureTargetTrait);
    }


}