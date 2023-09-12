import {DBService, ServiceTrait} from "./sdk/base_service";
import {log} from "./sdk/logger";
import {Day} from "./sdk/duration";
import {RaceReward, RaceRewardTrait} from "./unify/race_reward";
import {HDURaceReward} from "./school/hdu/hdu_race_reward";

interface RaceRewardLookupTrait{
    SchoolCode:string;
    StaffID:string;
}

export class RaceRewardService extends DBService implements ServiceTrait<RaceReward,RaceRewardLookupTrait>{
    constructor() {
        super("race-reward",{
            dataSourceID:"hdu-oracle-zxtd",
        },{
            ttl: 1*Day,
        });
    }

    lookup(conds: RaceRewardLookupTrait): RaceReward[] {
        let tx=this.getDB().startSession(HDURaceReward)
            .where("STAFFID=?",conds.StaffID);
        const raceRewardItems= tx.find() as HDURaceReward[];
        //getRuntime().debugBreakpoint(this,scoreItems);
        const unifyRaceRewardItems= raceRewardItems.map((item)=>{
            return new RaceReward({
                SchoolCode: "hdu",
                StaffID: item.StaffID,
                StaffName: item.StaffName,
                RaceName: item.RaceName,
                RaceLevel: item.RaceLevel,
                RewardLevel: item.RewardLevel,
                RewardDate: item.RewardDate,
            })
        });
        log(unifyRaceRewardItems);
        return unifyRaceRewardItems;
    }

    newUnifyModel():RaceReward{
        return new RaceReward({} as RaceRewardTrait);
    }
}