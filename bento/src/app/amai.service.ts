import { Injectable } from '@angular/core';
import 'rxjs/add/operator/map';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class AmaiService {

    //local API url
    url : string = "http://localhost:9000/";
    /*since i'm new at ts i don't exactly understatnd why
     * the constructor requires de http var on it
     */
    constructor(private http: HttpClient ) { }

    /* this functions returns the array of object that is given by the api
     */
    getProductsObservable(): Observable<any> {
        return this.http.get(this.url+"show");
    }

    deleteProduct( id: number ) {
        console.log("delete");
        this.http.delete(this.url+id).subscribe();
    }

    createProduct( code: string, price: number ) {
        console.log("create");
        this.http.get(this.url+"insert/"+code+"/"+price).subscribe();
    }
}
