
type EntityClassType<T extends SourceEntityTrait>={new(props:any):T};
export interface DBTrait {
    startSession<T extends SourceEntityTrait>(ctor:EntityClassType<T>):DBSessionTrait<T>;
}

export interface DBSessionTrait<T>{
    where(query:any,...args:any[]):DBSessionTrait<T>;
    find():T[];
}

export interface SourceEntityTrait {

}