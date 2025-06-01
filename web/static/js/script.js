function limparEtapasApos(etapaId) {
    const etapas = ['etapa-linhas', 'etapa-pecas', 'etapa-modelos', 'etapa-tecidos'];
    const index = etapas.indexOf(etapaId);
    if (index === -1) return;

    for (let i = index + 1; i < etapas.length; i++) {
        const el = document.getElementById(etapas[i]);
        if (el) el.innerHTML = '';
    }
}

document.addEventListener('click', function (e) {
    const target = e.target.closest('[data-etapa]');
    if (target) {
        const etapaId = target.getAttribute('data-etapa');
        limparEtapasApos(etapaId);
    }
});

