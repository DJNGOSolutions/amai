import { NgModule } from '@angular/core';
import { RouterModule, Routes, ExtraOptions } from '@angular/router';
import { CommonModule } from '@angular/common';

import { HomeComponent }        from './home/home.component';
import { UsersTableComponent }        from './users-table/users-table.component';
import { CreateUserComponent }        from './create-user/create-user.component';
import { CrudComponent }        from './crud/crud.component';
import { PanelComponent }   from './panel/panel.component';

// Pretty self explicatory we specify a route and it's component.
const routes: Routes = [
    { path: '', component: UsersTableComponent },
    { path: 'tools', component: CreateUserComponent },
    { path: 'panel', component: PanelComponent }
];

@NgModule({
    imports: [ RouterModule.forRoot(routes)],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}
