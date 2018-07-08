import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { Observable,of } from 'rxjs';
import { User } from './models/User';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json'
      /*'Authorization': 'my-auth-token'*/
  })
}; 

@Injectable()
export class AmaiService {

    //local API url
    url : string = "http://localhost:9000";
    //url : string = "";

    constructor(private http: HttpClient) { }

    /* this functions returns the array of object that is given by the api
     */

    createUser( name: string, age: number, gender: string, userEmail: string,
        aLevel: string, role: string ) {
        console.log("create");
        
        this.http.post(this.url+"/insert",
            {UserName:name, UserAge:age, UserGender:gender, UserEmail:userEmail,
                UserAcademicLevel: aLevel, UserRole: role})
            .subscribe(
            res => Response );
    }
     getUsers(): Observable<any> {
             return this.http.get(this.url+"/show");
    } 

    getGenders(): Observable<any>{
         return this.http.get(this.url+"/genders");
    }

    getLevels() : Observable<any>{
         return this.http.get(this.url+"/academics");
    }

    createSubject(name:string){
        this.http.post(this.url+"/create/subject/"+name,httpOptions).subscribe( res => Response );
    }

    createCategory(name:string){
        this.http.post(this.url+"/create/category/"+name,httpOptions).subscribe( res => Response );
    }

    createLevel(name:string){
        this.http.post(this.url+"/create/academic/"+name,httpOptions).subscribe( res => Response );
    }
}
