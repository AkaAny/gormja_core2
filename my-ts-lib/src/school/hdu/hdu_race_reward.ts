import {dbField, DBType} from "../../sdk/db_field";
import {SourceEntityTrait} from "../../sdk/db";

export class HDURaceReward implements SourceEntityTrait{
    @dbField("RECORDID",DBType.string)
    RecordID:string;
    @dbField("STAFFID",DBType.string)
    StaffID:string;
    @dbField("STAFFNAME",DBType.string)
    StaffName:string;
    @dbField("RACENAME",DBType.string)
    RaceName:string;
    @dbField("RACELEVEL",DBType.string)
    RaceLevel:string;
    @dbField("REWARDLEVEL",DBType.string)
    RewardLevel:string;
    @dbField("REWARDDATE",DBType.string)
    RewardDate:string;

    constructor(props:any) {
        this.RecordID=props.RecordID;
        this.StaffID=props.StaffID;
        this.StaffName=props.StaffName;
        this.RaceName=props.RaceName;
        this.RaceLevel=props.RaceLevel;
        this.RewardLevel=props.RewardLevel;
        this.RewardDate=props.RewardDate;
    }

    static tableName(): string {
        return "HDUHELP_VIEW_RACE_REWARD";
    }

}