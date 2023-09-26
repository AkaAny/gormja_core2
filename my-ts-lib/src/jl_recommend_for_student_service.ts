import {DBService, ServiceTrait} from "./sdk/base_service";
import {Hour} from "./sdk/duration";
import {log} from "./sdk/logger";
import {JLRecommendForStudent, JLRecommendForStudentTrait} from "./unify/jl_recommend_for_student";
import {JLRecommendForStudentItem} from "./center/jl_recommend_for_student_item";

interface JLRecommendForStudentLookupTrait{
    SchoolCode:string;
    StaffID:string;
}

export class JLRecommendForStudentService extends DBService implements ServiceTrait<JLRecommendForStudent, JLRecommendForStudentLookupTrait>{
    //ttl with manual refresh api
    constructor() {
        super("jl-recommend-for-student",{
            dataSourceID:'hdu-jl-recommend-output',
        },{
            ttl:1*Hour,
        });
    }

    lookup(params: JLRecommendForStudentLookupTrait): JLRecommendForStudent[] {
        let tx=this.getDB().startSession(JLRecommendForStudentItem)
            //.where("school_code=?",params.SchoolCode)
            .where("student_id=?",params.StaffID);
        log(params);
        const items= tx.find() as JLRecommendForStudentItem[];
        const results=items.map((item)=>{
            return new JLRecommendForStudent({
                SchoolCode: "hdu",
                StaffID: item.StaffID,
                JobID: item.JobID,
                Score: item.Score,
            })
        })
        log(results);
        return results;
    }

    newUnifyModel():JLRecommendForStudent{
        return new JLRecommendForStudent({} as JLRecommendForStudentTrait);
    }


}