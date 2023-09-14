import {DBService, ServiceTrait} from "./sdk/base_service";
import {Day} from "./sdk/duration";
import {log} from "./sdk/logger";
import {JLInternshipExperience, JLInternshipExperienceTrait} from "./unify/jl_internship_experience";
import {JLInternshipExperienceItem} from "./center/jl_internship_experience_item";

interface JLInternshipExperienceLookupTrait{
    SchoolCode:string;
    StaffID:string;
}

export class JLInternshipExperienceService extends DBService implements ServiceTrait<JLInternshipExperience, JLInternshipExperienceLookupTrait>{
    //ttl with manual refresh api
    constructor() {
        super("jl-internship-experience",{
            dataSourceID:'hdu-jl-student-attribute',
        },{
            ttl:1*Day,
        });
    }

    lookup(params: JLInternshipExperienceLookupTrait): JLInternshipExperience[] {
        let tx=this.getDB().startSession(JLInternshipExperienceItem)
            .where("school_code=?",params.SchoolCode)
            .where("staff_id=?",params.StaffID);
        log(params);
        const items= tx.find() as JLInternshipExperienceItem[];
        const results=items.map((item)=>{
            return new JLInternshipExperience({
                ID: item.ID,
                SchoolCode: item.SchoolCode,
                StaffID: item.StaffID,
                CompanyName: item.CompanyName,
                JobName: item.JobName,
                StartAt: item.StartAt,
                EndAt: item.EndAt,
                Description: item.Description,
            })
        })
        log(results);
        return results;
    }

    newUnifyModel():JLInternshipExperience{
        return new JLInternshipExperience({} as JLInternshipExperienceTrait);
    }


}