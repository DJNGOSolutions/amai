import { NgModule } from '@angular/core';
import { RouterModule, Routes, ExtraOptions } from '@angular/router';
import { CommonModule } from '@angular/common';

import { HomeComponent }        from './home/home.component';
import { CrudComponent }        from './crud/crud.component';

// Pretty self explicatory we specify a route and it's component.
const routes: Routes = [
    { path: '', component: HomeComponent },
    { path: 'crud', component: CrudComponent }
];

@NgModule({
    imports: [ RouterModule.forRoot(routes)],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}
