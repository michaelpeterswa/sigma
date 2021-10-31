import {writable} from "svelte/store";

class App {
  name: string;
  version: string;

  constructor(data:{name:string, version:string}) {
    this.name = data.name;
    this.version = data.version;
  }
}
export { App };
export const applist = writable<App[]>([]);