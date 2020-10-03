export default class{
    constructor(deps){
        this.tokenKey = deps.tokenKey
        this.proto = deps.proto
        this.client = new deps.proto.CategoryClient('http://localhost:9002', null, null)
    }

    index (){
        const req = new this.proto.IndexRequest()
        return this.client.index(req, {}).then(res=>{
            return res.getCategoriesList().map(category=>{
                return {
                    id: category.getId(),
                    name: category.getName(),
                }
            });
        })
    }
}