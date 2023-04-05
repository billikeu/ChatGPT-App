export namespace backend {
	
	export class AccountInfo {
	    chat_engine: string;
	    openai_api_key: string;
	    base_url: string;
	    openai_access_token: string;
	    newbing_cookies: string;
	    proxy: string;
	
	    static createFrom(source: any = {}) {
	        return new AccountInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chat_engine = source["chat_engine"];
	        this.openai_api_key = source["openai_api_key"];
	        this.base_url = source["base_url"];
	        this.openai_access_token = source["openai_access_token"];
	        this.newbing_cookies = source["newbing_cookies"];
	        this.proxy = source["proxy"];
	    }
	}
	export class AccountState {
	    account_info: AccountInfo;
	
	    static createFrom(source: any = {}) {
	        return new AccountState(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.account_info = this.convertValues(source["account_info"], AccountInfo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

