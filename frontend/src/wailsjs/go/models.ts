export namespace domain {
	
	export class NengaAtena {
	    id: number[];
	    year: number;
	    names: string[];
	    postcode: string;
	    address: string;
	    group_name: string;
	    kind: number;
	    is_sender: boolean;
	
	    static createFrom(source: any = {}) {
	        return new NengaAtena(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.year = source["year"];
	        this.names = source["names"];
	        this.postcode = source["postcode"];
	        this.address = source["address"];
	        this.group_name = source["group_name"];
	        this.kind = source["kind"];
	        this.is_sender = source["is_sender"];
	    }
	}

}

