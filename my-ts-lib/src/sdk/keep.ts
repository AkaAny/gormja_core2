export function forceKeep(...args:any[]){
    for (let arg in args){
        if(!arg){
            continue;
        }
        arg.toString(); //do something meaningless
    }
}