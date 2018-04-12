import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import { Products } from './Products';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/do';
import 'rxjs/add/operator/filter';
@Injectable()
export class AmaiService {

    //local API url
    url : string = "";
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
        this.http.delete(this.url+"delete/"+id).subscribe();
    }

    createProduct( code: string, price: number ) {
        console.log("create");
        this.http.post("insert",{Code:code, Price:price}).subscribe();
    }
}
