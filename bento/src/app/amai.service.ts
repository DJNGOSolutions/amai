import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable,of } from 'rxjs';
import { User } from './User';
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
         return this.http.get(this.url+"/show");
     } 

    deleteProduct( id: number ) {
        console.log("delete");
        this.http.delete(this.url+"/delete/"+id).subscribe();
    }

    createProduct( name: string, age: number ) {
        console.log("create");
        this.http.post("/insert",{NombreUsuario:name, EdadUsuario:age}).subscribe(
            res => Response );
    }
}
