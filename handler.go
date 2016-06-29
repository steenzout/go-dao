package dao

type TransactionFunc func(m Manager, ctx *Context) error

type TransactionFunc2 func(m Manager, ctx *Context) (interface{}, error)

func Process(m Manager, f TransactionFunc) error {
	ctx, err := m.StartTransaction()
	if err != nil {
		return err
	}
	defer m.EndTransaction(ctx)

	err = f(m, ctx)
	if err != nil {
		m.RollbackTransaction(ctx)
		return err
	}

	if err = m.CommitTransaction(ctx); err != nil {
		m.RollbackTransaction(ctx)
		return err
	}
	return nil
}

func Process2(m Manager, f TransactionFunc2) (interface{}, error) {
	ctx, err := m.StartTransaction()
	if err != nil {
		return nil, err
	}
	defer m.EndTransaction(ctx)

	obj, err := f(m, ctx)
	if err != nil {
		m.RollbackTransaction(ctx)
		return nil, err
	}

	if err = m.CommitTransaction(ctx); err != nil {
		m.RollbackTransaction(ctx)
		return nil, err
	}
	return obj, nil
}