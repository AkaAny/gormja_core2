import "reflect-metadata";
import {forceKeep} from "./keep";

export enum DBType {
    string = "string",
    int64 = "int64",
    bool = "bool",
    float64 = "float64",
    dateTime="dateTime",
}

const metaDataKey = "dbField";//Symbol("dbField");

interface DBFieldMetaData {
    column: string;
    dbType: DBType;
}

export function dbField(column: string, dbType: DBType, {
    isPrimaryKey = false,
}: {
    isPrimaryKey?: boolean,
}={}) {
    return Reflect.metadata(metaDataKey, {
        column: column,
        dbType: dbType,
        isPrimaryKey: isPrimaryKey,
    } as DBFieldMetaData);
}

export function getDBField(target: any, propertyKey: string) {
    return Reflect.getMetadata(metaDataKey, target, propertyKey);
}

forceKeep(dbField, getDBField);
