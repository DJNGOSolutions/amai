import { Injectable } from '@angular/core';
import 'rxjs/add/operator/map';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class AmaiService {

    url : string = "http://localhost:9000";
    constructor(private http: HttpClient ) { }

    getProductsObservable(): Observable<any> {
        return this.http.get(this.url);
  }
}
