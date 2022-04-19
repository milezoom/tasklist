import * as models from './models';

export interface go {
  "main": {
    "App": {
		GetTasks():Promise<string>
		StoreTask(arg1:string):Promise<string>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
