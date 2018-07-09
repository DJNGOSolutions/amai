import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { Observable,of } from 'rxjs';
import { User } from './models/User';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json',
    'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV4YW1wbGVAZW1haWwyLmNvbSIsIm5iZiI6MTUzMDk2NDgwMH0.YhaqEjwdSSQ04AnJIZ-Nu7ETjgsNL-jvA1lSLRb4Ol4'
  })
}; 

@Injectable()
export class AmaiService {

    //local API url
    //url : string = "http://localhost:9000";
    //url : string = "";
    ////server url
    url : string = "http::206.189.212.245:9000";

    constructor(private http: HttpClient) { }

    /* this functions returns the array of object that is given by the api
     */

    createUser( name: string, age: string, gender: number, userEmail: string,pass: string, aLevel: number, role: number ) {
        console.log("create");
        
        this.http.post(this.url+"/register",
            {UserName:name,UserPassword: pass, UserBirthday:age, IdGenderUser:gender, UserEmail:userEmail,
                IdAcademicLevelUser: aLevel, IdRoleUser: role},httpOptions)
            .subscribe(res => Response);
    }

    deleteUser( id : number ){
        this.http.delete(this.url+"/delete/"+id,httpOptions).subscribe( res => { console.log(res); });
    }

     getUsers(): Observable<any> {
             return this.http.get(this.url+"/show",httpOptions);
    } 

    getGenders(): Observable<any>{
         return this.http.get(this.url+"/genders",httpOptions);
    }

    getLevels() : Observable<any>{
         return this.http.get(this.url+"/academics",httpOptions);
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
