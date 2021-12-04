import { expect } from 'chai';
import {hello} from '../hello'

describe('hello!', () => { 
    it('works', () => { 
        expect(hello('akash')).to.equal('hello akash!')
    });
});

