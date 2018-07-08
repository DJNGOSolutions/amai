import { Component, OnInit } from '@angular/core';
import { AmaiService } from '../amai.service';
import { User } from '../models/User';
import { UsersTableComponent } from '../users-table/users-table.component';

@Component({
  selector: 'app-crud',
  templateUrl: './crud.component.html',
  styleUrls: ['./crud.component.css'],
})

export class CrudComponent implements OnInit {
    /*we created a class named Products as a kind of template that will
     *contain the same info as the JSON from the API returns
     */
    users : User[];

    constructor(private service : AmaiService ) {
    }

    onSubmit( id: number ){
        console.log("on submit");
        this.update();
    }

    refresh($event){
        this.update();
    }

    update(){
    }

    ngOnInit() {
        this.update();
    }

}
