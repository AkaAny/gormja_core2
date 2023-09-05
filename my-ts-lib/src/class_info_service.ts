import {DBService, ServiceTrait} from "./sdk/base_service";
import {log} from "./sdk/logger";
import {ClassInfo, ClassInfoTrait} from "./unify/class_info";
import {HDUClassInfo} from "./school/hdu/hdu_class_info";

interface ClassInfoLookupTrait{
    SchoolCode:string;
    CounselorStaffID:string;
    Grade:string;
}

export class ClassInfoService extends DBService implements ServiceTrait<ClassInfo,ClassInfoLookupTrait>{
    constructor() {
        super("class-info",{
            dataSourceID:'hdu-oracle',
        },false);
    }

    lookup(conds:ClassInfoLookupTrait):ClassInfo[]{
        log(conds);
        let tx= this.getDB().startSession(HDUClassInfo)
            .where("FDY_STAFFID=?",conds.CounselorStaffID);
        if(conds.Grade){
            tx=tx.where("GRADE=?",conds.Grade);
        }
        const classInfos= tx.find() as HDUClassInfo[];
        log(classInfos);
        const results=classInfos.map((item)=>{
            return new ClassInfo({
                SchoolCode:'hdu',
                Grade:item.Grade,
                ClassID: item.ClassID,
                CounselorStaffID: item.CounselorStaffID,
            });
        })
        log(results);
        return results;
    }

    newUnifyModel():ClassInfo{
        return new ClassInfo({} as ClassInfoTrait);
    }
}