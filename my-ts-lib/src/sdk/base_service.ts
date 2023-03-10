import {Duration} from "./duration";
import {DBTrait} from "./db";

export enum SourceType{
    DB="db",Http="http"
}


export class DBInitParam{
    dataSourceID:string;

    constructor(dataSourceID:string) {
        this.dataSourceID=dataSourceID;
    }

}

interface CacheTTLTrait{
    ttl:Duration; //millisecond
}

type CachePolicy=false | CacheTTLTrait;

export interface ServiceTrait<T,W>{
    lookup(params:W):T[];
    newUnifyModel():T;
}

export class BaseService{
    serviceID:string;
    sourceType:SourceType;
    ttl:Duration;

    constructor(serviceID:string,sourceType:SourceType,cachePolicy:CachePolicy) {
        this.serviceID=serviceID;
        this.sourceType=sourceType;
        if(cachePolicy===false){
            this.ttl=0;
            return;
        }
        this.ttl=cachePolicy.ttl;
    }
}

export class DBService extends BaseService{
    dbInitParam:DBInitParam;
    private db:any;
    constructor(serviceID:string,dbInitParam:DBInitParam,cachePolicy:CachePolicy) {
        super(serviceID,SourceType.DB,cachePolicy);
        this.dbInitParam=dbInitParam;
    }

    // @ts-ignore
    readonly init=(db:any)=> {
        if (!db) {
            throw "this should be called by runtime";
        }
        this.db = db;
    }

    getDB():DBTrait{
        return this.db as DBTrait;
    }
}
