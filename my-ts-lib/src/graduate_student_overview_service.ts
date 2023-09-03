import {DBService, ServiceTrait} from "./sdk/base_service";
import {log} from "./sdk/logger";
import {GradeOverview, GradeOverviewTrait} from "./unify/grade_student_overview";
import {HDUGradeOverview} from "./school/hdu/hdu_grade_student_overview";

interface GraduateStudentOverviewLookupTrait{
    SchoolCode:string;
    Grade:number;
}

export class GraduateStudentOverviewService extends DBService implements ServiceTrait<GradeOverview,GraduateStudentOverviewLookupTrait>{
    constructor() {
        super("graduate-student-overview",{
            dataSourceID:'hdu-oracle',
        },false);
    }

    lookup(conds:GraduateStudentOverviewLookupTrait):GradeOverview[]{
        log(conds);
        let tx= this.getDB().startSession(HDUGradeOverview)
            .where("GRADE=?",conds.Grade);
        const overViewItems= tx.find() as HDUGradeOverview[];
        log(overViewItems);
        const results=overViewItems.map((item)=>{
            return new GradeOverview({
                SchoolCode:'hdu',
                Grade:item.Grade,
                GraduateStudentCount:item.GraduateStudentCount,
            });
        })
        log(results);
        return results;
    }

    newUnifyModel():GradeOverview{
        return new GradeOverview({} as GradeOverviewTrait);
    }
}