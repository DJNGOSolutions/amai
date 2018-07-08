import { Component, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort } from '@angular/material';
import { UsersTableDataSource } from './users-table-datasource';
import { AmaiService } from '../amai.service';

@Component({
  selector: 'app-users-table',
  templateUrl: './users-table.component.html',
  styleUrls: ['./users-table.component.css']
})
export class UsersTableComponent implements OnInit {
  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;
  dataSource: UsersTableDataSource;

    constructor(private service : AmaiService ) {}

    /** Columns displayed in the table. */
    displayedColumns = ['Options','IdUser','UserRole','UserName',
        'UserGender', 'UserEmail','UserAcademicLevel','UserRate'];

    deleteUser(id:number){
        console.log( "delete: " + id );
        this.service.deleteUser(id);
    }

  ngOnInit() {
    this.dataSource = new UsersTableDataSource(this.paginator, this.sort,this.service);
  }
}
