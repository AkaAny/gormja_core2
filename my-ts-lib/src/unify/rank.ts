import {dbField, DBType} from "../sdk/db_field";
import {forceKeep} from "../sdk/keep";
import {log} from "../sdk/logger";

export interface RankTrait{
    SchoolCode:string;

    StaffID:string;

    GPA:number;

    Rank:number;
}

export class Rank implements RankTrait{
    @dbField("school_code",DBType.string,{
        isPrimaryKey:true,
    })
    SchoolCode:string;

    @dbField("staff_id",DBType.string,{
        isPrimaryKey:true,
    })
    StaffID:string;

    @dbField("gpa",DBType.float64)
    GPA:number;

    @dbField("rank",DBType.float64)
    Rank:number;

    constructor(props:RankTrait) {
        this.SchoolCode=props.SchoolCode;
        this.StaffID=props.StaffID;
        this.GPA=props.GPA;
        this.Rank=props.Rank;
    }

    static tableName():string{
        return "ranks";
    }
}

const RankInstance=new Rank({
    SchoolCode:'hdu',
    StaffID:"20113128",
    GPA: 0,
    Rank: 0,
});

log(Reflect.ownKeys(Rank.prototype));

forceKeep(Rank,RankInstance);