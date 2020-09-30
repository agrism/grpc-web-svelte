export default class{
    constructor(deps){
        this.tokenKey = deps.tokenKey
        this.proto = deps.proto
        this.client = new deps.proto.CategoryClient('http://localhost:9002', null, null)
    }

    async index (){
        const req = new this.proto.IndexRequest()
        const res = await this.client.index(req, {})
        return res

    }
}