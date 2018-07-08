import { Component, OnInit } from '@angular/core';
import { User } from '../models/User';
import { Category } from '../models/Category';
import { Subject } from '../models/Subjects';
import { AcademicLevel } from '../models/AcademicLevel';
import { Gender } from '../models/Gender';
import { AmaiService } from '../amai.service';
import { EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-create-user',
  templateUrl: './create-user.component.html',
  styleUrls: ['./create-user.component.css']
})
export class CreateUserComponent implements OnInit {
    user : User = {
        UserName: '',
        Age: 0,
        Gender: 1,
        UserEmail: '',
        AcademicLevel: 1,
        Role: 1
    }

    category : Category = {
        TopicName: ''
    }

    subject : Subject = {
        SubjectName: ''
    }

    level : AcademicLevel = {
        AcademicLevel: ''
    }

    categories : Category[];
    subjects : Subject[];
    levels : AcademicLevel[];
    genders : Gender[] ;

    constructor( private service: AmaiService) { }

    createUser( name: string, age: number, gender: string, userEmail: string,
        aLevel: string, role: string) {
        console.log("on create");
        console.log("vars",name,age,gender,userEmail,aLevel,role);
        console.log("this user",this.user);
        this.service.createUser(name,age,gender,userEmail,aLevel,role);
        //reset placeholder
        this.user = {
            UserName: '',
            Age: 0,
            Gender: 1,
            UserEmail: '',
            AcademicLevel: 1,
            Role: 1,
        };
        console.log(this.user);
        //this.created.emit('created');
    }

    createSubject(name:string){
        this.service.createSubject(name);
        this.subject.SubjectName = '';
    }

    createCategory(name:string){
        this.service.createCategory(name);
        this.category.TopicName = '';
    }

    createLevel(name:string){
        this.service.createLevel(name);
        this.level.AcademicLevel = '';
    }

  ngOnInit() {
      this.service.getGenders().subscribe( res => {
          this.genders = res;
          console.log( this.genders );
      }); 

      this.service.getLevels().subscribe( res => {
          this.levels = res;
      });

  }

}
