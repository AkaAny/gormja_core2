import {dbField, DBType} from "../sdk/db_field";
import {log} from "../sdk/logger";

export interface RaceRewardTrait{
    SchoolCode:string;
    StaffID:string;
    StaffName:string;
    RaceName:string;
    RaceLevel:string;
    RewardLevel:string;
    RewardDate:string;
}

export class RaceReward implements RaceRewardTrait{
    @dbField("school_code",DBType.string,{
        isPrimaryKey:true,
    })
    SchoolCode:string;

    @dbField("staff_id",DBType.string,{
        isPrimaryKey:true,
    })
    StaffID:string;

    @dbField("staff_name",DBType.string)
    StaffName:string;

    @dbField("race_name",DBType.string)
    RaceName:string;

    @dbField("race_level",DBType.string)
    RaceLevel:string;

    @dbField("reward_level",DBType.string)
    RewardLevel:string;

    @dbField("reward_date",DBType.string)
    RewardDate:string;

    constructor(props:RaceRewardTrait) {
        this.SchoolCode=props.SchoolCode;
        this.StaffID=props.StaffID;
        this.StaffName=props.StaffName;
        this.RaceName=props.RaceName;
        this.RaceLevel=props.RaceLevel;
        this.RewardLevel=props.RewardLevel;
        this.RewardDate=props.RewardDate;
    }

    static tableName():string{
        return "race_rewards";
    }
}

log(Reflect.ownKeys(RaceReward.prototype));