//import './style.css'
//import typescriptLogo from './typescript.svg'
//import { setupCounter } from './counter'
import {getRuntime} from "./sdk/runtime";
import {RankService} from "./rank_service";
import {ProfileService} from "./profile_service";
import {ScoreService} from "./score_service";
//import {dbField} from "./sdk/db_field";

Promise.all([
    getRuntime().registerService(ProfileService),
    getRuntime().registerService(ScoreService),
    getRuntime().registerService(RankService)
]).then((services)=>{
    getRuntime().debugBreakpoint("after register services",services);
    // const rankService=services[2];
    // rankService.lookup({
    //     // Grade:'2020',
    //     // UnitCode:'27',
    //     // MajorCode:'2703',
    //     StaffID: '20113128',
    // });
}).catch((err)=>{
    getRuntime().debugBreakpoint("register all services catch",err);
});


// document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
//   <div>
//     <a href="https://vitejs.dev" target="_blank">
//       <img src="/vite.svg" class="logo" alt="Vite logo" />
//     </a>
//     <a href="https://www.typescriptlang.org/" target="_blank">
//       <img src="${typescriptLogo}" class="logo vanilla" alt="TypeScript logo" />
//     </a>
//     <h1>Vite + TypeScript</h1>
//     <div class="card">
//       <button id="counter" type="button"></button>
//     </div>
//     <p class="read-the-docs">
//       Click on the Vite and TypeScript logos to learn more
//     </p>
//   </div>
// `
//
// setupCounter(document.querySelector<HTMLButtonElement>('#counter')!)
