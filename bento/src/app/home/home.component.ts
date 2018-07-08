import { Component, OnInit } from '@angular/core';
import { AmaiService } from '../amai.service';
import { User } from '../models/User';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
    /*we created a class named Products as a kind of template that will
     *contain the same info as the JSON from the API returns
     */
    users : User[];

    constructor(private service : AmaiService ) {
    }

    ngOnInit() {
    }

}
