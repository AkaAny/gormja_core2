import {BaseService} from "./base_service";

export type ClassType<T>={new():T}

export interface RuntimeTrait{
    registerService<T extends BaseService>(ctor: ClassType<T>):Promise<T>;
    getService<T extends BaseService>(ctor:ClassType<T>):T;
    debugBreakpoint(hint:string,...scopes:any[]):void;
}

export function getRuntime():RuntimeTrait{
    // @ts-ignore
    if(!Runtime){
        throw "need to run in runtime";
    }
    // @ts-ignore
    const pRuntime=Runtime; //use global obj to avoid registerService from being optimized
    return pRuntime as RuntimeTrait;
}